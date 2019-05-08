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
    margin: 50px 0 0 0;

    @media (min-width: 500px) {
        margin: 30px;
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
                <Header activeKey={process.env.PUBLIC_URL + '/about'} />

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
                            <P>
                                The standard Lorem Ipsum passage, used since the 1500s
                                <br />
                                "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
                            </P>

                            <P>
                                Section 1.10.32 of "de Finibus Bonorum et Malorum", written by Cicero in 45 BC
                                <br />
                                "Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur? Quis autem vel eum iure reprehenderit qui in ea voluptate velit esse quam nihil molestiae consequatur, vel illum qui dolorem eum fugiat quo voluptas nulla pariatur?"
                            </P>

                            <P>
                                1914 translation by H. Rackham
                                <br />
                                "But I must explain to you how all this mistaken idea of denouncing pleasure and praising pain was born and I will give you a complete account of the system, and expound the actual teachings of the great explorer of the truth, the master-builder of human happiness. No one rejects, dislikes, or avoids pleasure itself, because it is pleasure, but because those who do not know how to pursue pleasure rationally encounter consequences that are extremely painful. Nor again is there anyone who loves or pursues or desires to obtain pain of itself, because it is pain, but because occasionally circumstances occur in which toil and pain can procure him some great pleasure. To take a trivial example, which of us ever undertakes laborious physical exercise, except to obtain some advantage from it? But who has any right to find fault with a man who chooses to enjoy a pleasure that has no annoying consequences, or one who avoids a pain that produces no resultant pleasure?"
                            </P>
                        </Col>
                    </Row>
                </section>
            </C>
        );
    }
}

export default About;
