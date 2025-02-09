import os, subprocess, shlex, json

# Run a process and print it's output
def run_process(command):
    print("\nCommand:")
    print(command, "\n")
    env = os.environ.copy()
    process = subprocess.Popen(
        shlex.split(command), 
        stdout=subprocess.PIPE,
        universal_newlines=True,
        env=env
    )
    while True:
        output = process.stdout.readline()
        if output != '':
            print(output.strip())
        # Do something else
        return_code = process.poll()
        if return_code is not None:
            print(f"{'🟢' if return_code == 0 else '❗️'} CODE: {return_code}\n")
            # Process has finished, read rest of the output 
            for output in process.stdout.readlines():
                print(output.strip())
            return return_code

# Run a process syncronously and return result
def run_process_sync_in_background_return_result(command):
    env = os.environ.copy()
    p = subprocess.Popen(
        shlex.split(command),
        stdout=subprocess.PIPE,
        stderr=subprocess.STDOUT,
        env=env
    )
    (output, err) = p.communicate()
    p_status = p.wait()
    return str(output)

# Load a yaml file without any dependencies
def load_config(env, path_prefix = ""):
    file_path = f"{path_prefix}config/{env}.cfg.yml"
    # test if file exists
    if not os.path.exists(file_path):
        print(f"❌ Config file for {env} not found at {file_path}")
        return {}, False
    with open(file_path, 'r') as file:
        data = {}
        current_dict = data
        stack = []
        current_indent = 0
        for line in file:
            line = line.strip()
            if not line or line.startswith('#'):
                continue
        indent = len(line) - len(line.lstrip())
        key, value = line.split(':', 1)
        key = key.strip()
        value = value.strip()
        if indent > current_indent:
            stack.append(current_dict)
            current_dict[key] = {}
            current_dict = current_dict[key]
        elif indent < current_indent:
            for _ in range((current_indent - indent) // 2):
                current_dict = stack.pop()
            current_dict[key] = value
        else:
            current_dict[key] = value
        current_indent = indent
    print(f"✅ Loaded config for {env} from {file_path}")
    print(json.dumps(data, indent=2))
    return data, True

def write_deployment_status(region, status):
    # Read the existing contents of the file
    try:
        with open("deployment_status.txt", "r") as f:
            lines = f.readlines()
    except FileNotFoundError:
        lines = []
    # Update the status for the given region if it exists, or add a new entry if it doesn't
    updated = False
    with open("deployment_status.txt", "w") as f:
        for line in lines:
            if line.startswith(f"{region}:"):
                f.write(f"{region}: {status}\n")
                updated = True
            else:
                f.write(line)
        if not updated:
            f.write(f"{region}: {status}\n")
    print(f"✅ Updated deployment_status.txt for {region} to {status}")
