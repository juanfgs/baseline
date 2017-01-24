{{ define "body"}}
<h1>Sections</h1>
<div class="panel panel-default">
  <div class="panel-body">
<table class="table table-striped">
    <tr>
        <th>Id</th>
        <th>Name</th>
        <th>Icon</th>
        <th>Actions</th>
    </tr>
    {{range $key, $section := .SectionList}}
    <tr>
        <td>{{ $section.Id}}</td>
        <td>{{ $section.Name}}</td>
        <td><i class="fa fa-{{ $section.Icon}}"></i></td>
        <td></td>
    </tr>
    
    {{end }}

  </div>
</div>

{{ end }}
