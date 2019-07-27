import React, { Component } from "react";
import { Navbar, Nav } from "react-bootstrap";
import styled from 'styled-components';
import '../../fonts/stylesheet.css';


const NavbarBrand = styled(Navbar.Brand)`
    font-family: "cubic-millimeter";
    font-size: 1.7em !important;
    line-height: 0 !important;
`


class Header extends Component {

    constructor(props, context) {
        super(props, context);
        this.state = {
            activeKey: props.activeKey,
        }
    }

    render = () => {
        const docsURL = process.env.REACT_APP_DOCS_URL;
        return (
          <Navbar collapseOnSelect expand="lg" bg="dark" variant="dark">
            <NavbarBrand href={process.env.PUBLIC_URL + '/#/'}>efimeral</NavbarBrand>
            <Navbar.Toggle aria-controls="responsive-navbar-nav" />
            <Navbar.Collapse id="responsive-navbar-nav">
              <Nav className="mr-auto"></Nav>
              <Nav defaultActiveKey={this.state.activeKey}>
                <Nav.Link href={process.env.PUBLIC_URL + '/#/'}>Home</Nav.Link>
                <Nav.Link href={docsURL}>API Docs</Nav.Link>
                <Nav.Link href={process.env.PUBLIC_URL + '/#/about'}>About</Nav.Link>
              </Nav>
            </Navbar.Collapse>
          </Navbar>
        );
    }
}

export default Header;
