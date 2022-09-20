import {alpha} from '@mui/material/styles';
import {InputAdornment} from "@mui/material";
import TextField from "@mui/material/TextField";
import {useEffect, useState} from "react";
import SearchIcon from '@mui/icons-material/Search';

function FilterBar(props:{search:(term:string)=>void, sx:object}) {
    const [filterTerm, setFilterTerm] = useState("")

    useEffect(() => {
        // console.log("use effect: " + filterTerm)
        props.search(filterTerm)
    }, [filterTerm])

    return      <TextField
        id="input-with-icon-textfield"
        InputProps={{
            startAdornment: (
                <InputAdornment position="start">
                    <SearchIcon sx={{color: 'white'}}/>
                </InputAdornment>
            ),
            sx: {
                backgroundColor: alpha('#ffffff', 0.15),
                color: 'white',
                borderBlock: 'none'
            }
        }}
        placeholder="Filter..."
        variant="outlined"
        size="small"
        sx={{
            borderStyle: 'none',
            // ...{backgroundColor: alpha('#ffffff', 0.15),},
        ...props.sx
    }}
        onChange={e => setFilterTerm(e.target.value)}
    />
    // return <Search sx={props.sx}>
    //     <SearchIconWrapper>
    //         <SearchIcon />
    //     </SearchIconWrapper>
    //     <StyledInputBase
    //         placeholder="Filter…"
    //         inputProps={{ 'aria-label': 'search' }}
    //         onChange={(e) => {
    //             console.log("before set: " + e.target.value)
    //             setFilterTerm(e.target.value)
    //             console.log("after set: " + e.target.value)
    //         }}
    //     />
    // </Search>


    // return <Container>
    //     <TextField
    //         autoFocus
    //         margin="dense"
    //         id="search"
    //         type="text"
    //         fullWidth
    //         label="Filter"
    //         onChange={e => props.search(e.target.value)}
    //     />
    // </Container>
}

export default FilterBar