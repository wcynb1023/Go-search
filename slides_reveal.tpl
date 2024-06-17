{%- extends 'default.tpl' -%}

{% block input %}
    <div class="cell_input">
        {{ super() }}
    </div>
{% endblock input %}

{% block output %}
    <div class="cell_output">
        {{ super() }}
    </div>
{% endblock output %}
