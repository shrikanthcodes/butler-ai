from openai import OpenAI
from Credentials import OPENAI_API_KEY

# Define your prompt
prompt = """Give me a recipe with tomatoes, garlic and chicken"""


client = OpenAI(api_key = OPENAI_API_KEY)

completion = client.chat.completions.create(
  model="gpt-3.5-turbo",
  messages=[
    {"role": "user", "content": prompt}
  ]
)

print(completion.choices[0].message)
