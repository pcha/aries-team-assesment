import {alpha} from '@mui/material/styles';
import {InputAdornment} from "@mui/material";
import TextField from "@mui/material/TextField";
import {useEffect, useState} from "react";
import SearchIcon from '@mui/icons-material/Search';

/**
 * Component with a textfield to filter products
 *
 * @param props - {
 *     handleFilter: handler to call when the text is written
 *     sx: styles to apply to the component
 * }
 * @constructor
 */
function FilterBar(props:{handleFilter:(term:string)=>void, sx:object}) {
    const [filterTerm, setFilterTerm] = useState("")

    useEffect(() => {
        // console.log("use effect: " + filterTerm)
        props.handleFilter(filterTerm)
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
}

export default FilterBar