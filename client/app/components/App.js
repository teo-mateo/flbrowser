import React from 'react';

import {BrowserRouter as Router, Link, Route, Switch} from 'react-router-dom'
import FlbrowserNav from './nav/FlbrowserNav'
import WebAPI from '../util/WebAPI'
import RTorrentList from './parts/RTorrentList'

import BrowseFL from "./parts/BrowseFL"
let browserHistory = Router.browserHistory;


class App extends React.Component{
    constructor(){
        super();
        this.handleClick = this.handleClick.bind(this);
    }

    handleClick(event){
        console.log("from handler")
    }

    renderHome(){
        return(
            <div>
                <FlbrowserNav/>
				<p>FLBrowser client 2</p>
                <BrowseFL category={1} page={1}/>
            </div>

        )
    }

    renderActive(){
        return(
            <div>
                <FlbrowserNav />
                <p>FLBrowser client - RTorrent</p>
                <RTorrentList />
            </div>
        )
    }
	/*

    renderSettings(){
        return (
            <div>
                <BittraderNav/>
                <Settings/>
            </div>

        )
    }

    renderPositions(){
        return (
            <div>
                <BittraderNav/>
                <Positions />
            </div>
        )
    }

    renderSimulations(){
        return (
            <div>
                <BittraderNav/>
                <Simulations />
            </div>
        )
    }

    renderMarkets(){
        return(
            <div>
                <BittraderNav/>
                <Markets />
            </div>
        )
    }
	*/

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
                                    <FlbrowserNav/>
                                    <p>FLBrowser client 2</p>
                                    <BrowseFL category={1} page={1}/>
                                </div>                                
                            )
                        }} />
                        <Route path='/active' render={this.renderActive} />
                    </Switch>
                </Router>
                <button onClick={this.handleClick}>Button</button>
            </div>
        )
    }
}


module.exports = App;