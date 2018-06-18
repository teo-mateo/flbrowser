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

        this.loadFLTorrents = this.loadFLTorrents.bind(this);

    }

    componentDidMount(){
        this.loadFLTorrents();
    }

    componentDidUpdate(){
        //this.loadFLTorrents();
    }
    componentWillReceiveProps(){
        this.loadFLTorrents();
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