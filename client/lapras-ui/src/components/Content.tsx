import {Container} from "@mui/material";
import Header from "./header/Header";
import LoginForm from "./LoginForm";
import Products from "./products/Products";
import useWindowDimensions from "../hooks/useWindowDimensions";
import useSession from "../hooks/useSession";
import useProducts from "../hooks/useProducts";

/**
 * General Component containing all the others
 * @constructor
 */
function Content() {
    const session = useSession()
    const [products, setFilter, createProduct] = useProducts(session)
    const windowDimensions = useWindowDimensions()

    return <div>
        <Header session={session} handleFilter={setFilter} />
        <LoginForm session={session}/>
        {session.isLoggedIn ? <Container id="content" sx={{}}>
            <Products products={products} maxHeight={windowDimensions.height - 152}
                      handleCreate={createProduct}/>
        </Container> : ""}
    </div>
}

export default Content