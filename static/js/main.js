class Section extends React.Component {
    render(){
        var icon = "fa fa-" + this.props.data.icon + " fa-5x";
        return (
        <div className="iconbox">
          <div className="iconbox-icon">
                <i className={icon}></i>
          </div>
          <div className="featureinfo">
                <h4 className="text-center">{this.props.data.name}</h4>
                <p>
                {this.props.data.tagline}
            </p>
                <a className="btn btn-default btn-sm" href="#"  onClick={ () => this.props.setCurrentSection(this.props.data)} role="button">View »</a>
          </div>
        </div>

        );
    }
}

class FullSection extends React.Component {
    renderItem(item,idx){
        return (<Item data={item} order={idx} itemType={this.props.data.itemType} />);
    }

    render(){
        console.log(this.props.data);
        var className = "items";
        var idName = "items";
        var element = "ul";
        if(this.props.data.itemType == 'accordion'){
            className ="panel-group";
            idName ="accordion";
            element="div";
        }
        var converter = new showdown.Converter();
        var htmlContent = converter.makeHtml(this.props.data.content);
        
        const CustomTag = element;
        return(
                <div className="panel col-md-12" id="mainSection">
                   <h2>{this.props.data.name}</h2>
                <p dangerouslySetInnerHTML={{__html:htmlContent }}></p>
                <CustomTag className={className} id={idName} role="tablist">
                {this.props.data.items.map(function(entry,idx){
                    return  <div className={ this.props.data.itemType == 'accordion' ? "panel panel-default" : ""} key={idx}> {this.renderItem(entry,idx)}</div>;
                }, this)}
            </CustomTag>
                </div>
        );
    }

}

class Item extends React.Component {
    render(){
        switch(this.props.itemType){
        case 'accordion':
        return(
            <div>
                <div className="panel-heading" role="tab" id={"heading" + this.props.order}>
                      <h4 className="panel-title">
                <a role="button" data-toggle="collapse" data-parent="#accordion" href={"#collapse" + this.props.order} aria-expanded="false" aria-controls="collapseOne">
                          {this.props.data.name}
                      </a>
                </h4>
                </div>
                <div id={"collapse" +this.props.order} className="panel-collapse collapse out" role="tabpanel" aria-labelledby={"heading" +this.props.order}>
                    <div className="panel-body">
                            {this.props.data.description}
                    </div>
                </div>
                </div>
              );

            break;
        default:
            return(
                    <li><h4>{this.props.data.name}</h4> {this.props.data.description} </li>
            );
        }
    }

}

class CV extends React.Component {
    constructor(){
        super();
        this.state = {
        };
        this.setCurrentSection = this.setCurrentSection.bind(this);
    }

    render() {
        if (typeof this.state.sections != 'undefined'){
            var currentSection = "";
            if(typeof this.state.currentSection != 'undefined'){
                currentSection = <FullSection data={this.state.currentSection} />;
            }
            return (
                  <div className="CV">
                    <div className="col-md-3">
                    {this.state.sections.map(function(entry,idx){
                        return <div key={idx}> {this.renderSection(entry)}</div>;
                    }, this)}
                    </div>
                    <div className="col-md-8">
                                        {currentSection}
                    </div>
                  </div>
            );
        } else {
            return (<div className="CV" >Loading</div>);
        }
    }
    renderSection(section){
        return <Section data={section} setCurrentSection={this.setCurrentSection} />;
    }
    setCurrentSection(section){
        var new_state = this.state;
        new_state["currentSection"] = section;
        this.setState(new_state);
        console.log(this.state)
    }
    componentDidMount(){
        getJson("/data.json", this, "sections");
    };

}


ReactDOM.render(
  <CV />,
    document.getElementById('cv-content')
);

function getJson(url, obj, prop){
        var request = new XMLHttpRequest();
        request.open('GET', url, true);
        request.onload = function(  ) {
            if (request.status >= 200 && request.status < 400) {
                // Success!
                var newstate = obj.state;
                try {
                    newstate[prop] = JSON.parse(request.response);

                } catch(e) {
                    console.log(e);
                }
                obj.setState( newstate );
            } else {

                // We reached our target server, but it returned an error
            }
        }
    request.onerror = function() {
        return request.error;
    };
    request.send();
}
