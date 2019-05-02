import React, { Component } from 'react';
import styled from 'styled-components';
import { Image } from 'react-bootstrap'
import '../../fonts/stylesheet.css';
import ubuntu from './ubuntu.png';
import arch from './arch.png';


const PLogo = styled.p`
    font-size: 3em;
    font-family: "cubic-millimeter";
    color: white;
    line-height: 1em;
    margin: 0;
    margin-top: -0.45em;
    text-align: center;

    @media (min-width: 500px) {
      font-size: 7em;
    }
`

const PSymbol = styled.p`
    font-size: 15em;
    font-family: "cubic-millimeter";
    color: white;
    line-height: 1em;
    margin-top: -0.35em;
    text-align: center;

    @media (min-width: 500px) {
      font-size: 25em;
    }
`

const Linux = styled(Image)`
    width: 40%;
    height: auto;
    margin: auto;    
    display: block;
`

class Logo extends Component {
    render() {
      return (<PLogo>efimeral</PLogo>);
    }
}


class Symbol extends Component {
    render() {
      return (<PSymbol>e</PSymbol>);
    }
}

class Ubuntu extends Component {
    render() {
      return <Linux src={ubuntu} />;
    }
}

class Arch extends Component {
    render() {
      return <Linux src={arch} />;
    }
}


export { Logo, Symbol, Ubuntu, Arch };
