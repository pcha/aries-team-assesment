import {useCookies} from "react-cookie";
import {useEffect, useState} from "react";
import {ApiUrl} from "../etc/constants";

export type Session = {
    token: string,
    username: string,
    isLoggedIn: boolean,
    logIn: (username:string, password:string) => void,
    logOut: () => void,
    loginResultMessage: string
}

function useSession(): Session {
    const [cookies, setCookies, removeCookies] = useCookies(['token', 'username'])
    const [loggedIn, setLoggedIn] = useState(cookies.token != "") // If the user reload the site assume that it's logged and try to renew the token
    const [loginResultMessage, setLoginResultMessage] = useState("")

    const handleLoggedIn = () => setLoggedIn(true)
    const logIn = (username: string, password: string) => {
        setLoginResultMessage("")
        fetch(ApiUrl + "/users/login", {
            method: "POST",
            body: JSON.stringify({
                username: username,
                password: password
            })
        })
            .then(res => {
                switch (res.status) {
                    case 200:
                        res.json()
                            .then(body => {
                                setToken(body.token)
                            }).then(() => handleLoggedIn())
                        break
                    case 401:
                    case 500:
                    default:
                        res.json().then(body => setLoginResultMessage(body.error || "unknown error"))
                }
            }).catch(() => setLoginResultMessage("Connection error"))
    }

    const logOut = () => {
        removeCookies('token')
        removeCookies('username')
        setLoggedIn(false)
    }

    const setToken = (token: string) => setCookies('token', token);
    const setUsername = (username: string) => setCookies('username', username)

    const renewToken = () => {
        if (!loggedIn) return
        fetch(ApiUrl + "/users/token/renew", {
            method: 'POST',
            headers: {
                'Authorization': "Bearer " + cookies.token
            }
        }).then(res => {
            switch (res.status) {
                case 200:
                    res.json()
                        .then(body => {
                            setToken(body.token)
                            setUsername(body.claims.username)
                        })
                    setTimeout(renewToken, 13 * 60 * 1000) // renew token in 13 minutes
                    break
                case 401:
                    logOut()
                // TODO: 500?
            }
        })
    }

    useEffect(() => {
        renewToken();
    }, [loggedIn])


    return {
        token: cookies.token,
        username: cookies.username,
        isLoggedIn: loggedIn,
        logIn: logIn,
        logOut: logOut,
        loginResultMessage: loginResultMessage
    }
}

export default useSession