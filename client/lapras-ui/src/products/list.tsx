import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import {Alert} from "@mui/material";


// TODO type Product

function ProductsList(props:{products:any[], isSearch:boolean}) {
    const headerColor = ""

    return (props.products.length ? <TableContainer component={Paper} sx={{maxHeight: 500}}>
        <Table stickyHeader>
            <TableHead sx={{backgroundColor: "black"}}>
                <TableRow>
                    <TableCell sx={{backgroundColor: headerColor}}>ID</TableCell>
                    <TableCell sx={{backgroundColor: headerColor}}>Name</TableCell>
                    <TableCell sx={{backgroundColor: headerColor}}>Description</TableCell>
                </TableRow>
            </TableHead>
            <TableBody>
                {props.products.map((prod) => (
                    <TableRow key={prod.id}>
                        <TableCell>{prod.id}</TableCell>
                        <TableCell>{prod.name}</TableCell>
                        <TableCell>{prod.description}</TableCell>
                    </TableRow>
                ))}
            </TableBody>
        </Table>
    </TableContainer> : <Alert severity={"info"}>
        { props.isSearch?
            "There are no products matching the criteria, try a different one." :
            "There are no products yet, start adding one."}
    </Alert>)
}

export default ProductsList