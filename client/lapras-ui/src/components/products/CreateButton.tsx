import {Add} from "@mui/icons-material";
import {Fab} from "@mui/material";
import {useState} from "react";

/**
 * Component with the button to create a new Product
 * @param props - {
 *     handleShowForm: handler to call when the button is clicked
 * }
 * @constructor
 */
function CreateButton(props: { handleShowForm: () => void }) {
    const [exp, setExp] = useState(false)

    return (<Fab variant={exp ? "extended" : "circular"}
                 color="primary"
                 size={exp ? "large": "medium"}
                 sx={{
                     position: 'fixed',
                     bottom: 32,
                     right: 32,
                 }}
            onClick={props.handleShowForm}
                 onMouseOver={() => setExp(true)}
                 onMouseOut={() => setExp(false)}
        >{exp ? "Create " : ""}<Add /></Fab>
    )
}

export default CreateButton