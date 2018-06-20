import React from 'react';

import {BrowserRouter, Link, Route, Switch} from 'react-router-dom'
import FlbrowserNav from './nav/FlbrowserNav'
import WebAPI from '../util/WebAPI'
import RTorrentList from './parts/RTorrentList'
import CookieUtil from '../util/CookieUtil'
import BrowseFL from "./parts/BrowseFL"
import Login from "./parts/Login"

let browserHistory = BrowserRouter.history;

class App extends React.Component{
    constructor(){
        super();
        this.renderHome = this.renderHome.bind(this);
        this.handleLogin = this.handleLogin.bind(this);
        this.handleLogout = this.handleLogout.bind(this);
        this.renderActive = this.renderActive.bind(this);
        this.renderHome = this.renderHome.bind(this);
        this.getBasename = this.getBasename.bind(this);

        this.state = {
            loggedIn: false,
            category: 1, 
            page: 0
        };
    }

    getBasename(){
        if (window.location.host == "localhost:8081"){ //served locally with webpack
            return "";
        } else if (window.location.host == "localhost:8888"){ //served locally from go backend 
            return "/app";
        } else { //served remote
            return "/flbrowser/app";
        } 
    }

    componentDidMount(){
        var at = CookieUtil.GetAccessTokenFromCookie();

        if (at !== undefined){
            WebAPI.ping()
                .then( (response) => {
                    console.log("ping ok");
                    this.setState({ loggedIn: true });
                })
                .catch ( (error) => {
                    console.log("ping nok");
                    this.setState({ loggedIn: false });
                });
        }

        this.setState({
            loggedIn:  (at !== undefined)
        });
    }

    handleLogin(){
        var at = CookieUtil.GetAccessTokenFromCookie();
        this.setState({
            loggedIn:  (at !== undefined)
        });
    }
    handleLogout(){
        CookieUtil.RemoveAccessTokenCookie();
        this.setState({
            loggedIn: false
        });
    }

    renderHome(params_){

        let category=1;
        let page=0;
        if (params_.match.params.category && params_.match.params.page){
            category = params_.match.params.category;
            page = params_.match.params.page;
        }

        let content = ""
        if (this.state.loggedIn){
            content = (<BrowseFL 
                category={category} 
                page={page} />)
        } else {
            content = (<Login onLogin={this.handleLogin} />)
        }

        return(
            <div>
                <FlbrowserNav isLoggedIn={this.state.loggedIn} onLogout={this.handleLogout}/>
				<p>FLBrowser client 2</p>
                {content}
            </div>
        )
    }

    renderActive(){
        return(
            <div>
                <FlbrowserNav isLoggedIn={this.state.loggedIn}/>
                <p>FLBrowser client - RTorrent</p>
                <RTorrentList />
            </div>
        )
    }

    render(){
        let basename = this.getBasename();
        return (
            <div>   
                <BrowserRouter basename={basename}>
                    <Switch>
                        <Route path='/' exact render={this.renderHome} />
                        <Route path='/home' render={this.renderHome} />
                        <Route path='/browse/:category/:page' render={(params)=>{
                            return (
                                <div>
                                    <FlbrowserNav isLoggedIn={this.state.isLoggedIn}/>
                                    <p>FLBrowser client - browse</p>
                                    <BrowseFL   category={params.match.params.category} 
                                                page={params.match.params.page}
                                                />
                                </div>                                
                            )
                        }} />
                        <Route path='/active' render={this.renderActive} />
                    </Switch>
                </BrowserRouter>
            </div>
        )
    }
}


module.exports = App;