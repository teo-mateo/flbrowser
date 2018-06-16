import React from 'react';

import {BrowserRouter as Router, Link, Route, Switch} from 'react-router-dom'
import FlbrowserNav from './nav/FlbrowserNav'
import WebAPI from '../util/WebAPI'
import RTorrentList from './parts/RTorrentList'
import CookieUtil from '../util/CookieUtil'
import BrowseFL from "./parts/BrowseFL"
import Login from "./parts/Login"

let browserHistory = Router.browserHistory;


class App extends React.Component{
    constructor(){
        super();
        this.renderHome = this.renderHome.bind(this);
        this.handleLogin = this.handleLogin.bind(this);
        this.handleLogout = this.handleLogout.bind(this);
        this.renderActive = this.renderActive.bind(this);
        this.renderHome = this.renderHome.bind(this);

        var at = CookieUtil.GetAccessTokenFromCookie();
        this.state = {
            loggedIn:  (at !== undefined)
        }
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

    renderHome(){
        let content = ""
        if (this.state.loggedIn){
            content = (<BrowseFL category={1} page={1} />)
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
        return (
            <div>
                <Router history={browserHistory}>
                    <Switch>
                        <Route path='/' exact render={this.renderHome} />
                        <Route path='/home' render={this.renderHome} />
                        <Route path='/browse/:category/:page' render={(params)=>{
                            return (
                                <div>
                                    <FlbrowserNav isLoggedIn={this.state.isLoggedIn}/>
                                    <p>FLBrowser client 2</p>
                                    <BrowseFL category={1} page={1}/>
                                </div>                                
                            )
                        }} />
                        <Route path='/active' render={this.renderActive} />
                    </Switch>
                </Router>
            </div>
        )
    }
}


module.exports = App;