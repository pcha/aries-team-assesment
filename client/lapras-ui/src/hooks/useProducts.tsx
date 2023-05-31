import {useEffect, useState} from "react";
import {ApiUrl} from "../etc/constants";
import {Session} from "./useSession";

export type Product = {
    id: number,
    name: string,
    description: string
}

export type CreationResult = {
    success: boolean,
    message: string
}


/**
 * Hook to handle product list and creation. When an endpoint gets a http code 401, it handles log out
 *
 * @param session - Used to get token and handle logOut on 401
 *
 * @return [
 *  Products filtered list,
 *  Function to set the filter term ,
 *  Function that creates a new Product and return a promise to handle results
 * ]
 */
function useProducts (session: Session): [Product[], (filter:string)=>void, (name:string, description:string)=>Promise<CreationResult>] {
    const [products, setProducts] = useState<Product[]>([])
    const [filter, setFilter] = useState("")

    // get all the products matching with the filter, if it's set
    const fetchProducts = () => {
        const endpoint = ApiUrl + "/products" + (filter != "" ?  "?q=" + encodeURI(filter) : "")
        fetch(endpoint, {
            headers: {
                "Authorization": "Bearer " + session.token
            }
        })
            .then(res => {
                switch (res.status) {
                    case 200:
                        res.json()
                            .then(res => setProducts(res))
                        break
                    case 401:
                        session.logOut
                        break
                }
            })
    }

    // executes the fetch when the filter changes or the user logs in
    useEffect(() => {
        if (session.isLoggedIn) fetchProducts()
    }, [filter, session.isLoggedIn])

    // creates a product and returns a Promise to handle the request results
    const createProduct = (name: string, description: string) => {
        let data = {
            name: name,
            description: description
        }
        let result = {
            success: false,
            message: ""
        }
        return fetch(ApiUrl + "/products", {
            method: 'POST',
            headers: {
                Accept: 'application/form-data',
                'Content-Type': 'application/json',
                Authorization: "Bearer " + session.token
            },
            body: JSON.stringify(data),
        })
            .then(
                (res) => {
                    switch (res.status) {
                        case 201:
                            result.success = true
                            result.message = "Product Created!"
                            break
                        case 400:
                            return res.json().then(body => {
                                result.success = false
                                result.message = body.error || "Invalid data"
                            })
                        case 401:
                            session.logOut()
                            break
                        case 500:
                        default:
                            result.success = false
                            result.message = "Internal error. Try again later."

                    }
                },
                (err) => {
                    alert(err)
                    console.log(err)
                }
            ).then(() => fetchProducts()).then(()=>result)
    }

    return [products, setFilter, createProduct]
}

export default useProducts