import React, { Component } from 'react';
import styled from 'styled-components';
import { Row, Col, Modal, Button } from 'react-bootstrap'
import C from '../../components/Container';
import Header from '../../components/Header';
import { Logo, Symbol, Ubuntu, Arch } from '../../components/Logo';


const P = styled.p`
    font-size: 1em;
    color: white;
    text-align: center;
    margin: 0;

    @media (min-width: 500px) {
        font-size: 1.5em;
        margin: 30px;
    }
`

const Title = styled(Row)`
    margin-top: 60px;
`

const Distributions = styled(Row)`
    margin-top: 30px;
    @media (max-width: 500px) {
        margin-top: 0;
        & > div {
            margin-top: 40px;
        }
    }
`


class About extends Component {

    constructor(props, context) {
        super(props, context);
        this.state = {
            show: false,
        }
    }

    showModal = (e) => {
       e.preventDefault();
       this.setState({show: true});
    }

    closeModal = () => {
        this.setState({show: false});
    }

    render = () => {
        return (
            <C fluid="true">
                <Header />

                <section>    
                    <Row>
                        <Col md="12" sm="12">
                            <Logo />
                        </Col>
                    </Row>
                </section>

                <section>
                    <Row>
                        <Col md="12" sm="12">
                            <P>Choose a distribution</P>
                        </Col>
                    </Row>
                </section>
            </C>
        );
    }
}

export default Home;
