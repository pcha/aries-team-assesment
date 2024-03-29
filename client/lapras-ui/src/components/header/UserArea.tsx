import {Avatar, Box, Divider, IconButton, ListItemIcon, Menu, MenuItem, Tooltip} from "@mui/material";
import {Logout} from "@mui/icons-material";
import React, {useState} from "react";
import {Session} from "../../hooks/useSession";

/**
 * Component to show the User Area
 * @param props - {
 *     session: user session to handle
 * }
 * @constructor
 */
function UserArea(props: { session:Session }) {
    const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
    const open = Boolean(anchorEl);
    const handleClick = (event: React.MouseEvent<HTMLElement>) => {
        setAnchorEl(event.currentTarget);
    };
    const handleClose = () => {
        setAnchorEl(null);
    };
    return (props.session.isLoggedIn ? <div style={{flexGrow: 0.2}}>
        <Box sx={{display: 'flex', alignItems: 'center', textAlign: 'center', float: 'right'}}>
            <Tooltip title={"Logged as " + props.session.username} sx={{float: 'right'}}>
                <IconButton
                    onClick={handleClick}
                    size="small"
                    sx={{ml: 2}}
                    aria-controls={open ? 'account-menu' : undefined}
                    aria-haspopup="true"
                    aria-expanded={open ? 'true' : undefined}
                >
                    <Avatar sx={{width: 32, height: 32}}></Avatar>
                </IconButton>
            </Tooltip>
        </Box>
        <Menu
            anchorEl={anchorEl}
            id="account-menu"
            open={open}
            onClose={handleClose}
            onClick={handleClose}
            PaperProps={{
                elevation: 0,
                sx: {
                    overflow: 'visible',
                    filter: 'drop-shadow(0px 2px 8px rgba(0,0,0,0.32))',
                    mt: 1.5,
                    '& .MuiAvatar-root': {
                        width: 32,
                        height: 32,
                        ml: -0.5,
                        mr: 1,
                    },
                    '&:before': {
                        content: '""',
                        display: 'block',
                        position: 'absolute',
                        top: 0,
                        right: 14,
                        width: 10,
                        height: 10,
                        bgcolor: 'background.paper',
                        transform: 'translateY(-50%) rotate(45deg)',
                        zIndex: 0,
                    },
                },
            }}
            transformOrigin={{horizontal: 'right', vertical: 'top'}}
            anchorOrigin={{horizontal: 'right', vertical: 'bottom'}}
        >
            <MenuItem>
                <Avatar/> {props.session.username}
            </MenuItem>
            <Divider/>
            <MenuItem onClick={props.session.logOut}>
                <ListItemIcon>
                    <Logout fontSize="small"/>
                </ListItemIcon>
                Logout
            </MenuItem>
        </Menu></div> : null)
   }


export default UserArea