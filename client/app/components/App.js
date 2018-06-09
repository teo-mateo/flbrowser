import React from 'react';
import BittraderNav from './nav/BittraderNav'
import Home from './parts/home/Home'
import Positions from './parts/positions/Positions'
import Markets from './parts/markets/Markets'
import Settings from './parts/settings/Settings'
import Simulations from './parts/simulations/Simulations'

import {BrowserRouter as Router, Link, Route, Switch} from 'react-router-dom'

class App extends React.Component{
    constructor(){
        super();
    }

    renderHome(){
        return(
            <div>
                <BittraderNav/>
                <Home />
            </div>

        )
    }

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

    render(){
        return (
            <div>
                <Router>
                    <Switch>
                        <Route path='/' exact render={this.renderHome} />
                        <Route path='/home' render={this.renderHome} />
                        <Route path='/settings' render={this.renderSettings} />
                        <Route path='/positions' render={this.renderPositions} />
                        <Route path='/simulations' render={this.renderSimulations} />
                        <Route path='/markets' render={this.renderMarkets} />
                    </Switch>
                </Router>
            </div>
        )
    }
}


module.exports = App;