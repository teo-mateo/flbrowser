import Axios from 'axios'

var ApiRoot = "http://localhost:8888";
var api = function(url){
    return ApiRoot + url
}

module.exports = {

    getCategories : function(){
        return Axios.get(api("/categories"));
    },
    getRtrTorrents : function(){
        return Axios.get(api("/torrents/rtr"))
    }
}