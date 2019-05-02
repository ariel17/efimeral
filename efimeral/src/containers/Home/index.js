import React, { Component } from 'react';
import styled from 'styled-components';
import { Container, Row, Col, Modal, Button } from 'react-bootstrap'
import Header from '../../components/Header';
import { Logo, Symbol, Ubuntu, Arch } from '../../components/Logo';


const C = styled(Container)`
    background-color: #282c34;
    height: 100%;
`

const P = styled.p`
    font-size: 1.5em;
    color: white;
    text-align: center;
    margin: 30px;
`

const Distributions = styled.section`
    margin: 80px;
`


class Home extends Component {

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
                            <Symbol />
                            <Logo />
                        </Col>
                    </Row>
                </section>

                <Distributions>
                    <Row>
                        <Col md="12" sm="12">
                            <P>Choose a distribution</P>
                        </Col>
                    </Row>
                    <Row>
                        <Col md="6" sm="12">
                            <a href="/ubuntu" onClick={this.showModal}><Ubuntu /></a>
                        </Col>
                        <Col md="6" sm="12">
                            <a href="/arch" onClick={this.showModal}><Arch /></a>
                        </Col>
                    </Row>
                    <Modal show={this.state.show} onHide={this.closeModal}>
                        <Modal.Header closeButton>
                            <Modal.Title>Efimeral is WIP</Modal.Title>
                        </Modal.Header>
                        <Modal.Body>Sorry, it's not ready yet :(</Modal.Body>
                        <Modal.Footer>
                            <Button variant="primary" onClick={this.closeModal}>Close</Button>
                        </Modal.Footer>
                    </Modal>
                </Distributions>
            </C>
        );
    }
}

export default Home;
