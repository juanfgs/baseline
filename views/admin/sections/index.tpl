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
        <td>
            <div class="btn-group">
            <a href="/admin/sections/delete/{{ $section.Id }}" class="btn btn-sm btn-default"><i class="fa fa-close"></i> Delete</a>
            <a href="/admin/sections/edit/{{ $section.Id }}" class="btn btn-sm btn-default"><i class="fa fa-edit"></i> Edit</a>
            </div>
        </td>
    </tr>
    {{end }}

</table>    
  </div>
</div>

{{ end }}
