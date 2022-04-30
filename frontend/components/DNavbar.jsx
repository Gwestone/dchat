import React from 'react'
import { Nav, Button, Navbar, Container, NavDropdown } from 'react-bootstrap';
import Link from 'next/link'

export default function DNavbar(){
    return(
        <Navbar expand="lg" bg="dark" variant="dark">
            <Container>
                <Navbar.Brand href="#home">Docker chat</Navbar.Brand>
                    <Nav className="me-auto">
                        <Nav.Link href="#features">Features</Nav.Link>
                    </Nav>
                    <Nav>
                        <Nav.Link href="/user/register">register</Nav.Link>
                        <Nav.Link href="/user/login">login</Nav.Link>
                    </Nav>
            </Container>
        </Navbar>
    )
}