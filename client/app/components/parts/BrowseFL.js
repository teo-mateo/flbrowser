import React from 'react'
import CategoryList from './CategoryList'
class BrowseFL extends React.Component{

    constructor(props){
        super(props);
    }

    render(){
        return(
            <CategoryList />
        );
    }

}

module.exports = BrowseFL