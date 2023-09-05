from md2pdf.core import md2pdf
from datetime import date
from jinja2 import Environment, PackageLoader, meta, Template

# Apply templates
def apply_template(templateName, variables):
    with open(f"store/templates/{templateName}") as file:
        template = Template(file.read())
        return template.render(**variables)
    
def get_variables(templateName):
    env = Environment(loader=PackageLoader('store', 'templates'))
    template_source = env.loader.get_source(env, templateName)
    parsed_content = env.parse(template_source)
    return meta.find_undeclared_variables(parsed_content)

def pdf_from_template(templateName, variables, outputFile):
    content = apply_template(templateName, variables)
    md2pdf(outputFile, md_content=content)

print(get_variables("sfaf.md"))

pdf_from_template("sfaf.md", {
    "data": date.today().strftime("%d/%m/%Y"),
    "nome": "Michele",
    "cognome": "Mastrogiovanni"
}, "test.pdf")

