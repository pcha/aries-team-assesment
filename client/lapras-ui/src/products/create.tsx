import Button from '@mui/material/Button';
import {useEffect, useState} from "react";
import CreateForm from "./createForm";
import CreateButton from "./createButton";

function ProductCreator(props:{create: (name: string, description: string) => void}) {
    const [open, setOpen] = useState(false)

    const handleOpen = () => {
        setOpen(true)
    }

    const handleClose = () => {
        setOpen(false)
    }

    const handleCreate = (name:string, description: string) => {
        props.create(name, description)
        setOpen(false)
    }

    return(
        <div>
           <CreateButton handleShowForm={handleOpen} />
            <CreateForm open={open} handleClose={handleClose} handleCreate={handleCreate} />
        </div>
    )
}

export default ProductCreator