import React from 'react'
import CategoryLink from './CategoryLink'
import WebAPI from '../../util/WebAPI'

class CategoryList extends React.Component{

    constructor(props){
        super(props);
        this.state = {
            categories: []
        };

        WebAPI.getCategories()
        .then((result) =>{
            console.log(result);
            this.setState({
                categories: result.data
            })
        })
        .catch((error)=>{
            console.log(error)
        });
    }

    render(){
        return (
            <div>
                {
                    this.state.categories.map((v,i)=><span key={v.id}><CategoryLink category={v} /> &nbsp;&nbsp;</span>   )
                }
            </div>
        )
    }
}


module.exports = CategoryList