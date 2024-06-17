{% extends 'slides_reveal.tpl' %}

{% block header %}
    <style>
        /* 自定义背景色 */
        .reveal {
            background-color: #F0F0F0;
        }

        /* 自定义字体样式 */
        .reveal .slides .slide {
            font-family: Arial, sans-serif;
        }
    </style>
    {{ super() }}
{% endblock header %}
