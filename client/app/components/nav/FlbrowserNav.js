import React from 'react'
import {Navbar, NavbarBrand, Nav, NavItem, NavLink } from 'reactstrap';

class FlbrowserNav extends React.Component{
    constructor(){
        super();
        this.state = {};
    }

    render(){
        return(
            <Navbar color="faded">
                <Nav className="ml-auto" navbar>
                    <NavItem>
                        <NavLink href="/home">Home</NavLink>
                    </NavItem>
                </Nav>
                <Nav>
                    <NavItem>
                        <NavLink href="/browse">Browse FL</NavLink>
                    </NavItem>
                </Nav>
                <Nav>
                    <NavItem>
                        <NavLink href="/active">Active</NavLink>
                    </NavItem>
                </Nav>
            </Navbar>
        )
    }
}

module.exports=FlbrowserNav;