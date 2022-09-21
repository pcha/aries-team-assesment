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

function useProducts (session: Session): [Product[], (filter:string)=>void, (name:string, description:string)=>Promise<CreationResult>] {
    const [products, setProducts] = useState<Product[]>([])
    const [filter, setFilter] = useState("")

    const fetchProducts = () => {
        const endpoint = ApiUrl + (filter == "" ? "/products" : "/products/search/?q=" + encodeURI(filter))
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
                    case 500:
                    default:
                        // TODO Show MUI alert
                        res.json().then((body) => alert(body.error))
                }
            })
    }

    useEffect(fetchProducts, [filter])

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