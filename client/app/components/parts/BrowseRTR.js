import React from 'react'
import WebAPI from '../../util/WebAPI'

class BrowseRTR extends React.Component{
    constructor(props){
        super(props)
        this.state = {
            rtrTorrents: []
        };
    }

    render(){
        <div>RTorrentList component</div>
    }

    componentDidMount(){
        WebAPI.getRtrTorrents()
            .then((response) => {
                console.log(response.data);
                this.setState({
                    rtrTorrents: response.data
                });
            })
            .catch((error) =>{
                console.log(error);
            })
    }
}

module.exports = BrowseRTR