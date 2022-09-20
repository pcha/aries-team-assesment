import './App.css'
import Products from "./products";
import {Container} from "@mui/material";
import Header from "./header";
import {useEffect, useState} from "react";
import LoginForm from "./LoginForm";
import {useCookies} from "react-cookie";
import {ApiUrl} from "./constants";

export default App

function App() {
    const [cookies, setCookies, removeCookies] = useCookies(['token'])
    const [loggedIn, setLoggedIn] = useState(cookies.token != "")
    const [filterTerm, setFilterTerm] = useState("")

    const handleLoggedIn = () => {
        setLoggedIn(true)
    }
    const logOut = () => {
        removeCookies('token')
        setLoggedIn(false)
    }

    const setToken = (token:string) => setCookies('token', token);

    const renewToken = () => {
        fetch(ApiUrl + "/users/token/renew", {
            method: 'POST',
            headers: {
                'Authorization': "Bearer " + cookies.token
            }
        }).then(res => {
            switch (res.status) {
                case 200:
                    res.json()
                        .then(body => setToken(body.token))
                    break
                case 401:
                    logOut()
            }
        })
    }

    useEffect(() => {
        if (loggedIn) {
            renewToken();
        }
    })

    return (
        <div className="App" style={{overflow: "hidden"}}>
            <Header loggedIn={loggedIn} handleLogOut={() => setLoggedIn(false)} handleSearch={setFilterTerm}/>
            <LoginForm show={!loggedIn} handleLogIn={handleLoggedIn} setToken={setToken}/>
            {loggedIn ? <Container id="content" sx={{}}>
                <Products apiURL={import.meta.env.API_URL} apiToken={cookies.token} logOut={logOut} filterTerm={filterTerm}/>
            </Container>: ""}
        </div>
    );
}
