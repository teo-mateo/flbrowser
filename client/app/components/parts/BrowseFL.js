import React from 'react'
import CategoryList from './CategoryList'
import WebAPI from '../../util/WebAPI'
import { Table } from 'reactstrap'

class BrowseFL extends React.Component{

    constructor(props){
        super(props);
        this.state = {
            torrents: [], 
            rtorrents: []
        }

        this.loadTorrents = this.loadTorrents.bind(this);
        this.loadFLTorrents = this.loadFLTorrents.bind(this);
        this.loadRTRTorrents = this.loadRTRTorrents.bind(this);
    }



    componentDidMount(){
        this.loadTorrents();
    }

    componentWillReceiveProps(){
        this.loadTorrents();
    }

    loadTorrents(){
        this.loadFLTorrents();
        this.loadRTRTorrents();
    }

    loadFLTorrents(){
        WebAPI.getFlTorrents(this.props.category, this.props.page)
            .then((response) => {
                console.log('--getFlTorrents.then--');
                console.log(response);
                this.setState({torrents:response.data});
                //console.log(this.state);
            })
            .catch ((error) => {
                console.log('--getFlTorrents.error--');
                console.log(error);
            });
    }

    loadRTRTorrents(){
        WebAPI.getRtrTorrents()
            .then((response) =>{
                console.log('--getRtrTorrents.then--');
                console.log(response);
                this.setState({rtorrents:response.data});
            })
            .catch((error) => {
                console.log('--getRtrTorrents.error--');
                console.log(error);
            })
    }

    render(){

        return(
            <div className="marg5px">
                <CategoryList onNavigateTo={this.props.onNavigateTo}/>
                <Table responsive dark>
                    <thead>
                        <tr>
                            <th>#</th>
                            <th>Name</th>
                            <th>Size</th>
                            <th>Downloaded</th>
                            <th>Seeds/Leechs</th>
                        </tr>
                    </thead>
                    <tbody>
                        {this.state.torrents.map((v,i) =>
                             (<tr key={i}>
                                <th scope="row"> 
                                    <a href="#" onClick={(e) => {
                                        e.preventDefault();
                                        var id = e.target.innerText;
                                        WebAPI.downloadTorrent(id);
                                        alert("Torrent " + id + " was sent for download.");
                                    }}> {v.id} </a> </th>
                                <td> {v.name} </td>
                                <td> {v.size}</td>
                                <td> {v.timesdownloaded}</td>
                                <td> L/S {v.leechers} / {v.seeders}</td>
                            </tr>)
                        )}
                    </tbody>
                </Table>
            </div>
            

        );
    }

}

module.exports = BrowseFL