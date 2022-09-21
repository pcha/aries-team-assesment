import {AppBar, Box, Toolbar, Typography} from "@mui/material";
import FilterBar from "./FilterBar";
import UserArea from "./UserArea";
import {Session} from "../../hooks/useSession";

/**
 * Header (App Bar) component with the filter bar and the logged in area
 *
 * @param props - {
 *     session: Used lo handle log in / log out,
 *     handleFilter: handler to call when the user writes a text in the filter
 * }
 * @constructor
 */
function Header(props: { session: Session, handleFilter: (term: string) => void }) {
    return <Box sx={{flexGrow: 1}}>
        <AppBar position="fixed">
            <Toolbar>
                <Typography variant="h6" component="div" sx={{flexGrow: 0.2}}>
                    LAPRAS
                </Typography>
                {props.session.isLoggedIn ?
                    <FilterBar handleFilter={props.handleFilter} sx={{flexGrow: 0.6}}/> : ""}

                    <UserArea session={props.session} />


            </Toolbar>
        </AppBar>
        <Toolbar></Toolbar>
    </Box>

}

export default Header