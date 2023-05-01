
import openai

def summarize_cpp_code(code):
    # Prepare the prompt for GPT-3.5
    prompt = "Code: " + code + "\n\nSummarize the functionality of the code:"

    # Set up OpenAI API credentials
    openai.api_key =   API_KEY # Replace with your OpenAI API keyPI

    # Generate the summary using GPT-3.5
    response = openai.Completion.create(
        engine='davinci',
        prompt=prompt,
        max_tokens=100,
        temperature=0.5,
        n=1,
        stop=None,
        timeout=10
    )

    # Extract the summary from the API response
    summary = response.choices[0].text.strip()

    return summary

# Example usage
cpp_code = """
#include <iostream>

int main() {
    std::cout << "Hello, World!" << std::endl;
    return 0;
}
"""

summary = summarize_cpp_code(cpp_code)
print("Code Summary:", summary)
