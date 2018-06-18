import Axios from 'axios'
import CookieUtil from './CookieUtil'

var ApiRoot = "http://localhost:8888";
var api = function(url){
    return ApiRoot + url
}

module.exports = {

    getSecureOptionsObject: function(){
        let accessToken = CookieUtil.GetAccessTokenFromCookie();
        return {
            headers: {
                'FLAccessToken': accessToken
            }
        };
    },
    ping: function(){
        return Axios.get(api("/ping"), this.getSecureOptionsObject() );
    },
    login: function(obj){
        return Axios.post(api("/login2"), obj);
    },
    getCategories : function(){
        
        return Axios.get(api("/categories"), this.getSecureOptionsObject() );
    },
    getRtrTorrents : function(){
        return Axios.get(api("/torrents/rtr"), this.getSecureOptionsObject() );
    },
    getFlTorrents: function(category, page){
        return Axios.get(api("/torrents/fl/"+category+"/"+page), this.getSecureOptionsObject() );
    },
    downloadTorrent: function(id){
        return Axios.post(api("/torrents/fl/"+id+"/download"), {}, this.getSecureOptionsObject() );
    }


}