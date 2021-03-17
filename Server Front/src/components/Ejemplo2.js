import React from 'react'
import ReadDom from 'react-dom'

const Server = "http://localhost:7000";

class Ejemplo2 extends React.Component{
    constructor(props){
        super(props);
        this.state = {value: ""}
        
    }

    componentDidMount() {
        this.get_data_server()
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
                    <h1>Segundo ejemplo de prueba </h1>
                </div>
                <div className="card" >

                    <h5 className="card-header">Datos obtenidos del servidor</h5>
                    <div className="card-body">
                        <h5 className="card-title">{this.state.value}</h5>
                    </div>
                </div>
            </div>
        )
    }
}

ReadDom.render(
    <Ejemplo2 />,
    document.getElementById('root')
);

export default Ejemplo2;