import React from 'react'
import ReadDom from 'react-dom'

const Server = "http://localhost:7000";

class Ejemplo1 extends React.Component{
    constructor(props){
        super(props);
        console.log("Primer elemento")
        this.state = {value: ""}
    }

    get_data_server = async() => {
        const response = await fetch(`${Server}/getHello`);
        const data = await response.json();
        this.setState({value: data["respuesta"]});
    }
    render(){
        return(
            <div  style={{textAlign: 'center', width: '70%',paddingLeft: '15%'}}>
                <div className="header" >
                    <h1>Primer ejemplo de prueba </h1>
                </div>
                <div className="card" >

                    <h5 className="card-header">Datos obtenidos del servidor</h5>
                    <div className="card-body">
                        <h5 className="card-title">{this.state.value}</h5>
                        <button onClick={this.get_data_server} className="btn btn-primary">Obtener</button>
                    </div>
                </div>
            </div>
        )
    }
}

ReadDom.render(
    <Ejemplo1 />,
    document.getElementById('root')
);

export default Ejemplo1;