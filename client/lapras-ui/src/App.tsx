import './App.css'
import Products from "./products";
import {Container} from "@mui/material";
import Header from "./header";
import {useEffect, useState} from "react";
import LoginForm from "./LoginForm";
import {useCookies} from "react-cookie";
import {ApiUrl} from "./constants";
import useWindowDimensions from "./useWindowDimensions";

export default App

function App() {
    const [cookies, setCookies, removeCookies] = useCookies(['token', 'username'])
    const [loggedIn, setLoggedIn] = useState(cookies.token != "")
    const [filterTerm, setFilterTerm] = useState("")

    const handleLoggedIn = () => {
        setLoggedIn(true)
    }
    const logOut = () => {
        removeCookies('token')
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
            }
        })
    }

    useEffect(() => {
        renewToken();
    }, [loggedIn])

    const windowDimensions = useWindowDimensions()

    return (
        <div className="App" style={{overflow: "hidden"}}>
            <Header loggedIn={loggedIn} handleLogOut={logOut} handleSearch={setFilterTerm} username={cookies.username}/>
            <LoginForm show={!loggedIn} handleLogIn={handleLoggedIn} setToken={setToken}/>
            {loggedIn ? <Container id="content" sx={{}}>
                <Products apiURL={import.meta.env.API_URL} apiToken={cookies.token} logOut={logOut}
                          filterTerm={filterTerm} maxHeight={windowDimensions.height - 152}/>
            </Container> : ""}
        </div>
    );
}
