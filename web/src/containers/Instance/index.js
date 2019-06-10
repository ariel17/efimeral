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
            loading: true
        }
    }

    componentDidMount = () => {
        const apiURL = process.env.REACT_APP_API_URL;
        fetch(apiURL+"/sessions", {
            method: "POST",
            body: JSON.stringify({
                "distribution": "ubuntu",
                "tag": "19.04"
            })
        })
        .then(response => response.json()) 
        .then(data => {
            console.log(data);
            this.setState({
                loading: false
            })
        });
    }

    render = () => {
        return (
            <C fluid="true">
                <Header activeKey="" />

                <section>
                    <Row>
                        <Col md="12" sm="12">
                            {this.state.loading && 
                                <h3>LOADING ...</h3>
                            }
                            {!this.state.loading && 
                                <h3>SUCCESSS :)</h3>
                            }

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
