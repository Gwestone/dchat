import React, {useState} from 'react'
import {Nav, Button, Navbar, Container, NavDropdown, Form, InputGroup} from 'react-bootstrap';
import Link from 'next/link'
import "bootstrap/dist/css/bootstrap.min.css"
import DNavbar from "../../components/DNavbar";
import Error from "../../components/Error";

export default function login(){

    let [username, setUsername] = useState("")
    let [password, setPassword] = useState("")
    let [error, setError] = useState(null)

    const registerUser = async (e) => {
        e.preventDefault()
        try {
            const res = await fetch('http://dchat.org:81/register', {
                body: JSON.stringify({
                    username: username,
                    password: password
                }),
                headers: {
                    'Content-Type': 'application/json'
                },
                method: 'POST'
            })
        }catch (e){
            console.log("error")
            setError(e)
        }

    }

    const changeUsername = (e)=>{
        setUsername(e.target.value)
        e.preventDefault()
    }

    const changePassword = (e)=>{
        setPassword(e.target.value)
        e.preventDefault()
    }

    return(
        <>
            <DNavbar />
            <Container >
                <Form onSubmit={registerUser}>
                    <InputGroup className={"d-grid gap-2"}>
                        <input className={"username"} type={"text"} name={"username"} autoComplete={"username"} onChange={changeUsername} value={username} required/>
                        <input className={"password"} type={"password"} name={"password"} autoComplete={"new-password"} onChange={changePassword} value={password} required/>
                        <Button type={"submit"} >Register</Button>
                    </InputGroup>
                </Form>
                { error ? <Error err={error}/> : null }
            </Container>
        </>
    )
}