import React from 'react'
import { InputGroup, InputGroupAddon, InputGroupText, Input, Button } from 'reactstrap';
import WebAPI from '../../util/WebAPI'
import CookieUtil from '../../util/CookieUtil'

class Login extends React.Component{
    constructor(props){
        super(props);
        this.state={
            username:"",
            password:""
        }

        this.handleUsernameChange = this.handleUsernameChange.bind(this);
        this.handlePasswordChange = this.handlePasswordChange.bind(this);
        this.handleSubmitClick = this.handleSubmitClick.bind(this);
    } 

    handleUsernameChange(event){
        this.setState({username: event.target.value})
    }

    handlePasswordChange(event){
        this.setState({password: event.target.value})
    }

    handleSubmitClick(event){
        if(this.state.username === "" || this.state.password === ""){
            return
        }
        WebAPI.login(this.state)
            .then((response) => {
                console.log("-in then");
                console.log(response); 
                CookieUtil.SetAccessTokenCookie(response.data.accessToken, new Date(Date.now()+30*24*3600*1000)) // 30 days 
                this.props.onLogin()
            })
            .catch((error) => {
                console.log("-in catch");
                console.log(error);
                console.log(error.response);
            });
    }

    render(){

        return (<div className="centered">
            <InputGroup className="marg5px">
                <InputGroupAddon addonType="prepend">@</InputGroupAddon>
                <Input placeholder="username" value={this.state.username} onChange={this.handleUsernameChange} />
            </InputGroup>
            <InputGroup className="marg5px">
                <InputGroupAddon addonType="prepend">@</InputGroupAddon>
                  <Input placeholder="password" type="password" value={this.state.password} onChange={this.handlePasswordChange} />
                <InputGroupAddon addonType="append">
                  <Button onClick={this.handleSubmitClick}>Login</Button>
                </InputGroupAddon>
            </InputGroup>
            <br />
            <br />
        </div>)
    }
}

module.exports = Login