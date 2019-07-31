import React, { Component } from 'react';
import styled from 'styled-components';
import { Row, Col } from 'react-bootstrap'
import C from '../../components/Container';
import Header from '../../components/Header';
import { Logo } from '../../components/Logo';


const P = styled.p`
    font-size: 1em;
    color: white;
    text-align: left;
    margin: 20px;

    @media (min-width: 500px) {
        margin: 50px;
    }
`

const StyledSection = styled.section`
    margin: 20px 20px 0 20px;
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
                <Header activeKey={process.env.PUBLIC_URL + '/#/about'} />

                <section>    
                    <Row>
                        <Col md="12" sm="12">
                            <Logo />
                        </Col>
                    </Row>
                </section>

                <StyledSection>
                    <Row>
                        <Col md="12" sm="12">
                            <P>
                                <h2>What?</h2>
                                Efimeral creates Linux containers accessible
                                through any web browser and disposables in a
                                short period of time. Ubuntu and ArchLinux are
                                the available images for public usage and free,
                                customized with a lot of development
                                tools and languajes.<br />
                            </P>
                        </Col>
                        <Col md="6" sm="12">
                            <P>
                                LANGUAGES/MANAGERS:<br />
                                - SDKMan<br />
                                - Python / Pip<br />
                                - Java (OpenJDK) / Maven<br />
                                - Golang / dep<br />
                                - NodeJS / NPM / Yarn<br />
                                - PHP / Composer<br />
                            </P>
                        </Col>
                        <Col md="6" sm="12">
                            <P>
                                TOOLS:<br />
                                - curl<br />
                                - jq<br />
                                - httpie<br />
                                - telnet<br />
                                - htop<br />
                                - ab<br />
                                - vim<br />
                                - nano<br />
                                - git<br />
                                - grunt<br />
                                - gulp<br />
                                - ant<br />
                                - loadtest<br />
                            </P>
                        </Col>
                        <Col md="12" sm="12">
                            <P>
                                <h2>Why?</h2>
                                For fun, of course! :)<br />
                                <br />
                                Some years ago I was working in IT configuring
                                web servers and networking. I am not a network
                                expert, so I used to perform checks on my work
                                using an external point (my house computer
                                usually), but sometimes it was not an option
                                either. "How could I use a temporary point in
                                the Internet to do some tests?" I though, and
                                this concept came to my mind. Soon I realized
                                that it could be a good fit for others
                                searching for temporal access to a server, like
                                when you use a mobile device travelling to
                                work.
                            </P>
                            <P>
                                <h2>Who?</h2>
                                I am <a href="http://ariel17.com.ar">Ariel Gerardo Ríos</a> and I work as a
                                Software Engineer. I do stuff in my free time.
                            </P>
                        </Col>
                    </Row>
                </StyledSection>
            </C>
        );
    }
}

export default About;
