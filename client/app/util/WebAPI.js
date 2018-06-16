import Axios from 'axios'

var ApiRoot = "http://localhost:8888";
var api = function(url){
    return ApiRoot + url
}

module.exports = {

    login: function(obj){
        return Axios.post(api("/login2"), obj);
    },
    getCategories : function(){
        return Axios.get(api("/categories"));
    },
    getRtrTorrents : function(){
        return Axios.get(api("/torrents/rtr"))
    }

}