import React, { Component } from 'react';
import styled from 'styled-components';
import { Row, Col } from 'react-bootstrap'
import C from '../../components/Container';
import Header from '../../components/Header';
import Loader from '../../components/Loader';


const Iframe = styled.iframe`
    height: 93vh;
    width: 100vw;
`


class Instance extends Component {

    constructor(props, context) {
        super(props, context);
        this.state = {
            distribution: props.distribution,
            tag: props.tag,
            loading: true,
            url: null
        }
    }

    componentDidMount = () => {
        const apiURL = process.env.REACT_APP_API_URL;
        fetch(apiURL+"/sessions", {
            method: "POST",
            body: JSON.stringify({
                "distribution": this.state.distribution,
                "tag": this.state.tag
            })
        })
        .then(response => response.json()) 
        .then(data => {
            console.log(data);
            this.setState({
                loading: false,
                url: data.url
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
                                <Loader loading={this.state.loading} />
                            }
                            {!this.state.loading && 
                                <Iframe src={this.state.url} title="Container instance" frameBorder="0" allowfullscreen></Iframe>
                            }
                        </Col>
                    </Row>
                </section>
            </C>
        );
    }
}

export default Instance;
