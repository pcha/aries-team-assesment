import {AppBar, Box, Toolbar, Typography} from "@mui/material";
import Button from "@mui/material/Button";
import FilterBar from "./products/FilterBar";


function Header(props:{loggedIn:boolean, handleLogOut: () => void, handleSearch: (term:string) => void}) {
    return <Box sx={{flexGrow: 1}}>
        <AppBar position="fixed">
            <Toolbar>
                <Typography variant="h6" component="div" sx={{flexGrow: 0.2}}>
                    LAPRAS
                </Typography>
                {props.loggedIn ?
                <FilterBar search={props.handleSearch} sx={{flexGrow: 0.6}} />:""}
                {props.loggedIn ?
                    <div style={{flexGrow: 0.2}}>
                        <Button sx={{float: 'right'}} color="inherit" variant="text" onClick={props.handleLogOut}>{"Log Out"}</Button></div> :
                    ""}

            </Toolbar>
        </AppBar>
    <Toolbar></Toolbar>
    </Box>

}

export default Header