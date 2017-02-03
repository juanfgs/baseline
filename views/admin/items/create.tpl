{{ define "body" }}
<h1>Add Item</h1>
<form method="POST">
    <div class="form-group">
        <label for="Name">Name</label>
        <input type="text" class="form-control" name="Name" placeholder="Enter Item Name"/>
    </div>
    
    <div class="form-group">
        <label for="Value">Value</label>
        <input type="text" class="form-control" name="Value" placeholder="Enter Item Value"/>
    </div>

    <div class="form-group">
        <label for="Section">Section</label>
        
        <select name="SectionId">
        {{range $key, $section := .SectionList}}
            <option value="{{ $section.Id}}">
                {{ $section.Name}} - {{ $section.Icon}}                
            </option>
            {{end }}
        </select>
    </div>

    <button class="btn btn-primary" >Add</button>
</form>

{{ end }}
