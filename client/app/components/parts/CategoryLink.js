import React from 'react'
import {Link} from 'react-router-dom'

class CategoryLink extends React.Component{

    constructor(props){
        super(props);
    }

    render(){
        let url = '/browse/'+this.props.category.id + '/1'
        return(
            <Link to={url}>
            <span>{this.props.category.id} | {this.props.category.name}</span>
            </Link>
        )
    }
}

module.exports = CategoryLink