import {useCookies} from "react-cookie";
import {useEffect, useState} from "react";
import {ApiUrl} from "../etc/constants";

// Contain session related states and function
export type Session = {
    token: string,
    username: string,
    isLoggedIn: boolean,
    logIn: (username: string, password: string) => void,
    logOut: () => void,
    loginResultMessage: string
}

// Hook to handle user session
function useSession(): Session {
    const [firstRender, setFirstRender] = useState(true)
    const [cookies, setCookie, removeCookie] = useCookies(['token', 'username'])
    const [loggedIn, setLoggedIn] = useState(false) // If the user reload the site assume that it's logged and try to renew the token
    const [loginResultMessage, setLoginResultMessage] = useState("")

    // state representing the cookies values. This in needed to expose the values typed, otherwise typing prevent the value update
    const [username, setUsernameState] = useState<string>(cookies.username)
    const [token, setTokenState] =useState<string>(cookies.token)
    // use effect listening cookies changes and updating corresponding states
    useEffect(() => setUsernameState(cookies.username), [cookies.username])
    useEffect(() => setTokenState(cookies.token), [cookies.token])
    // use effect to handle logged in, it depends on token state because is when any component can access to the token.
    // Otherwise, it could result in eventual 401 responses.
    useEffect(() => setLoggedIn(!!token), [token])


    // wrapper to set the token in the cookies
    const setToken = (tkn: string) => setCookie('token', tkn);
    // wrapper to set the username in the cookies
    const setUsername = (username: string) => setCookie('username', username)

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
                            })
                        break
                    case 401:
                    case 500:
                    default:
                        res.json().then(body => setLoginResultMessage(body.error || "unknown error"))
                }
            }).catch(() => setLoginResultMessage("Connection error"))
    }

    const logOut = () => {
        removeCookie('token')
        removeCookie('username')
        setLoggedIn(false)
    }

    // Call the renovation api endpoint and set the new token, if fails it'll log out
    const renewToken = () => {
        if (!loggedIn || !cookies.token) {
            logOut()
            return
        }
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
                    break
                case 401:
                case 500:
                default:
                    logOut()
            }
        })
    }

    // To be executed only on the first render
    useEffect(() => {
        // mark the first render as executed
        setFirstRender(false)
        // If there is a token on a first render I don't know when it'll expire, so it forces a renovation
        renewToken()
    }, [])

    // When the token change, it schedules a renovation before expiration
    useEffect(() => {
        // In the first render the token in renewed to avoid expiration, so the renovation will be already scheduled.
        // Also, if the token is empty, doesn't make sense schedule a renovation
        if (!firstRender && cookies.token) {
            setTimeout(renewToken, 13 * 60 * 1000) // renew token in 13 minutes
        }
    }, [cookies.token])


    return {
        token: token,
        username: username,
        isLoggedIn: loggedIn,
        logIn: logIn, //executes log In
        logOut: logOut, //executer log out
        loginResultMessage: loginResultMessage
    }
}

export default useSession