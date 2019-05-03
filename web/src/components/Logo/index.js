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
    text-align: center;

    @media (min-width: 500px) {
      font-size: 7em;
    }
`

const PSymbol = styled.p`
    font-size: 12em;
    font-family: "cubic-millimeter";
    color: white;
    line-height: 0.9;
    margin: 0;
    text-align: center;

    @media (min-width: 500px) {
        font-size: 20em;
        line-height: 0.7;
    }
`

const Linux = styled(Image)`
    width: 75%;
    height: auto;
    margin: auto;    
    display: block;

    @media (min-width: 500px) {
        width: 40%;
    }
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
