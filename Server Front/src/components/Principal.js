import React from 'react'
import ReadDom from 'react-dom'


class Principal extends React.Component{
   
    render(){
        return(
            <div style={{backgroundColor: 'blueviolet'}}>
                <img src="https://images.pexels.com/photos/210186/pexels-photo-210186.jpeg?auto=compress&cs=tinysrgb&dpr=3&h=750&w=1260" alt="Nature" width="100%"/>
            </div>
        )
    }
}

ReadDom.render(
    <Principal />,
    document.getElementById('root')
);

export default Principal;