import React from 'react';
import './App.css';
import {
  Row,
  Col,
  Button,
  Container,
  Form
} from 'react-bootstrap';

var values = {
  level: "1",
  algorithm: "profundidad"
}


function App() {
  return (
    <div className="App">
      <header className="App-header">
        <h2>Seleccione el nivel y el algoritmo.</h2>
        <Container>
          <Form>
            <Form.Row>
              <Col>
                <Form.Group controlId="levels">
                  <Form.Label>Nivel</Form.Label>
                  <Form.Control as="select" value={values.level}>
                    <option>1</option>
                    <option>2</option>
                    <option>3</option>
                    <option>4</option>
                  </Form.Control>
                </Form.Group>
              </Col>
              <Col>
                <Form.Group controlId="algorithm">
                  <Form.Label>Algoritmo</Form.Label>
                  <Form.Control as="select" value={values.algorithm}>
                    <option value="profundidad">Preferente por profundidad</option>
                    <option value="anchura">Preferente por anchura</option>
                    <option value="iterativa">Profundidad iterativa</option>
                  </Form.Control>
                </Form.Group>
              </Col>
              <Col>
                <Button variant="primary" type="submit">
                  Ejecutar
                </Button>
              </Col>
            </Form.Row>
          </Form>

        </Container>
      </header>
    </div>
  );
}

function test(eventKey, event) {
  console.log(event)
}

export default App;
