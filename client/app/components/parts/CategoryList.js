import React from 'react'
//import { withRouter } from 'react-router-dom'
import CategoryLink from './CategoryLink'
import WebAPI from '../../util/WebAPI'
import { Button, ButtonGroup } from 'reactstrap';
import { withRouter } from 'react-router';

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
            });
        })
        .catch((error)=>{
            console.log(error)
        });
    }

    render(){
        return (
            <ButtonGroup>
                {this.state.categories.map((v,i)=>(
                        <Button key={v.id} 
                            onClick={(event) =>{
                                this.props.history.push( '/browse/'+v.id + '/1')
                        }}>{v.name}</Button>
                    ))}
            </ButtonGroup>
        )
    }
}


module.exports = withRouter(CategoryList)