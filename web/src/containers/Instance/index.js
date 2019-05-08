import React, { Component } from 'react';
import styled from 'styled-components';
import { Row, Col } from 'react-bootstrap'
import C from '../../components/Container';
import Header from '../../components/Header';


const P = styled.p`
    font-size: 1em;
    color: white;
    text-align: left;
    margin: 0;

    @media (min-width: 500px) {
        margin: 30px;
    }
`


class Instance extends Component {

    constructor(props, context) {
        super(props, context);
        this.state = {
            distro: props.distro,
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
                <Header activeKey="" />

                <section>
                    <Row>
                        <Col md="12" sm="12">
                            <P>
                                // TODO: Here goes the instance iframe
                            </P>

                        </Col>
                    </Row>
                </section>
            </C>
        );
    }
}

export default Instance;
