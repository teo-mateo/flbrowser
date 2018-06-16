import Cookies from 'universal-cookie'

const ucookie = new Cookies();
const accessTokenCookieName = "flbrowser-access-token"

module.exports = {
    GetAccessTokenFromCookie: function(){
        let atc = ucookie.get(accessTokenCookieName);
        console.log("Access token cookie: ", atc);
        return atc;
    },
    SetAccessTokenCookie: function(accessToken, expires){
        ucookie.set(accessTokenCookieName, accessToken, {
            path: '/',
            expires: expires
        });
    }, 
    RemoveAccessTokenCookie: function(){
        ucookie.remove(accessTokenCookieName);
    }

}