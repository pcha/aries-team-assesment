import {AppBar, Box, Toolbar, Typography} from "@mui/material";
import FilterBar from "./products/FilterBar";
import UserArea from "./UserArea";


function Header(props: { loggedIn: boolean, handleLogOut: () => void, handleSearch: (term: string) => void, username: string }) {
    return <Box sx={{flexGrow: 1}}>
        <AppBar position="fixed">
            <Toolbar>
                <Typography variant="h6" component="div" sx={{flexGrow: 0.2}}>
                    LAPRAS
                </Typography>
                {props.loggedIn ?
                    <FilterBar search={props.handleSearch} sx={{flexGrow: 0.6}}/> : ""}
                {props.loggedIn ?
                    <UserArea username={props.username} handleLogOut={props.handleLogOut}/>
                    :
                    ""}

            </Toolbar>
        </AppBar>
        <Toolbar></Toolbar>
    </Box>

}

export default Header