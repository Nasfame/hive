# CoopHive

CoopHive is a two-sided marketplace for computational resources. It enables users to run computational workloads in a permissionless protocol, where anyone can get paid to connect
their compute nodes to the network and run jobs.


# Getting started

Welcome to the prerelease series of CoopHive.

## CoopHive Networks

CoopHive operates on a decentralized infrastructure across multiple EVM-compatible networks.

### Add to Metamask

<details>


<summary>Aurora Testnet <i>(Deprecated)</i></summary>


<pre>

Network name: Aurora Testnet


RPC URL: http://aurora.co-ophive.network:8545


Chain ID: 1337


Currency symbol: ETH


Block Explorer URL: (leave blank)

</pre>


</details>

<details>

<summary>Halcyon Testnet</summary>

<pre>
Network name: Halcyon Testnet

RPC URL: http://halcyon.co-ophive.network:8545

Chain ID: 1337

Currency symbol: ETH

Block Explorer URL: (leave blank)
</pre>

</details>

<details>
<summary>Sepolia Testnet</summary>

- [x] Visit https://chainlist.org/chain/11155111
- [x] Add to Metamask

</details>

<details>
<summary>FVM Calibration Testnet</summary>

- [x] Visit https://chainlist.org/chain/314159
- [x] Add to Metamask

</details>




> The testnet has a base currency of ETH and you will also get HIVE to pay for jobs (and nodes to stake).

### Faucets

To obtain funds, please visit the below faucets:

1. Aurora Faucet <i>(Deprecated)</i>: [Click Here](http://faucet.co-ophive.network:8080)
2. Halcyon Faucet: [Click Here](http://halcyon-faucet.co-ophive.network:8085)
3. Sepolia Faucet:
   - [Hive Faucet](http://faucet.co-ophive.network:8081)
   - [Eth Faucet](https://sepoliafaucet.com)
4. Calibration Faucet: [Click Here](http://faucet.co-ophive.network:8082)

[//]: # (3. Sepolia Faucet: [Click Here]&#40;http://faucet.co-ophive.network:8081&#41;)

> The faucet will give you both ETH (to pay for gas) and HIVE (to stake and pay for jobs).


[//]: # (### Quick start on Sepolia Testnet)

[//]: # ()

[//]: # (- [ ] Add to [Metamask]&#40;https://chainlist.org/chain/11155111&#41;)

[//]: # (- [ ] Claim ETH drips from [Faucet]&#40;https://www.alchemy.com/faucets/ethereum-sepolia&#41;)

[//]: # (- [ ] Claim HIVE drips from [Coophive Faucet]&#40;http://faucet.co-ophive.network:8081&#41;)

## Install Hive Client

Download the latest release of CoopHive for your platform. Both the amd64/x86_64 and arm64 variants of macOS and Linux
are supported. (If you are on Apple Silicon, you'll want arm64).

To check your version use `which hive`. It it's an old version, run `rm <path>` to remove that path, and then reinstall the newest version.

The commands below will automatically detect your OS and processor architecture and download the correct CoopHive build for your machine.

### On Command Line

[//]: # (1. Detect your operating system and set it as $OSNAME)

[//]: # (2. Detect your machine's architecture and set it as $OSARCH)

[//]: # (3. Download the latest production build)

[//]: # (4. Check the version)

[//]: # (5. Install `hive`)

#### Install only hive

```bash
curl -sSf https://raw.githubusercontent.com/CoopHive/hive/main/install.sh | sh -s -- hive
```

#### For Resource providers and mediators: Install hive + bacalhau

```bash
curl -sSf https://raw.githubusercontent.com/CoopHive/hive/main/install.sh | sh -s -- all
```

[//]: # (<details> )

[//]: # (<summary>Installation script for Linux and MacOS</summary>)

[//]: # ()

[//]: # (```bash)

[//]: # (OSARCH=$&#40;uname -m | awk '{if &#40;$0 ~ /arm64|aarch64/&#41; print "arm64"; else if &#40;$0 ~ /x86_64|amd64/&#41; print "amd64"; else print "unsupported_arch"}'&#41; && export OSARCH)

[//]: # (echo $OSARCH)

[//]: # (OSNAME=$&#40;uname -s | awk '{if &#40;$1 == "Darwin"&#41; print "darwin"; else if &#40;$1 == "Linux"&#41; print "linux"; else print "unsupported_os"}'&#41; && export OSNAME;)

[//]: # (echo $OSNAME)

[//]: # (version=v0.10.0)

[//]: # (curl -sSL -o hive https://github.com/CoopHive/hive/releases/download/$version/hive-$OSNAME-$OSARCH)

[//]: # (chmod +x hive)

[//]: # (./hive version)

[//]: # ()

[//]: # (sudo mv hive /usr/local/bin/hive)

[//]: # (```)

[//]: # ()

[//]: # (</details>)

### Manual [GUI]

1. Go to https://github.com/CoopHive/hive/releases/latest
2. Download the binary for your system.

### With Go 1.21+

`go install github.com/CoopHive/hive@latest`

## Run a job

First, make sure your Web3 private key is in the environment.

```
export WEB3_PRIVATE_KEY=<your private key>
```

Alternatively, arrange for the key to be in your environment in a more secure way that doesn't get written to your shell history.

## Checkout our standard modules

<details>
<summary>Update: You can now switch from the default network</summary>

> using `--network <network>` or by setting the env var `export NETWORK=`

- `hive run cowsay:v0.1.2 -i Message="CoopHive" --network halcyon`
- `hive run cowsay:v0.1.2 -i Message="CoopHive" --network calibration`
- `hive run cowsay:v0.1.2 -i Message="CoopHive" --network aurora`
- `hive run cowsay:v0.1.2 -i Message="CoopHive" --network sepolia`
- `export NETWORK=sepolia && hive run cowsay:v0.1.2 -i Message="CoopHive"`

</details>

### Cowsay Module

Hello CoopHive, now with cows.

```shell
hive run cowsay:v0.1.2 -i Message="Hiro is building Coophive"
```

[//]: # (FIXME: the video uses the input var Prompt not Message. Upload fresh video)
https://github.com/CoopHive/hive/assets/24226219/d1768527-57d8-47cc-8643-4bb57d190769


<details>

<summary> Preview Run </summary>

<pre>
hive run cowsay:v0.1.2 -i Message="Hiro is building CoopHive" --network sepolia

  ___  __    __  ____  _  _  __  _  _  ____
 / __)/  \  /  \(  _ \/ )( \(  )/ )( \(  __)
( (__(  O )(  O )) __/) __ ( )( \ \/ / ) _)
 \___)\__/  \__/(__)  \_)(_/(__) \__/ (____) 0.17.1

  Decentralized Compute Network  https://coophive.network


●∙∙ CoopHive submitting job2024-03-08T07:21:28+05:30 INF ../../Volumes/N/CoopHive/hive/pkg/module/utils.go:148 > updating cached git repo=/Users/hiro/.coophive/sepolia/jc/repos/CoopHive/coophive-module-cowsay
🌟  CoopHive submitting job
∙∙∙ Job submitted. Negotiating deal...Deal QmSirREjZ9icDawWiR8EMQn8UaVMFySrocgfdGstCrjiGR is matched
∙●∙ Job submitted. Negotiating deal...2024/03/08 07:21:29 Deal QmSirREjZ9icDawWiR8EMQn8UaVMFySrocgfdGstCrjiGR is agreed
🤝  Job submitted. Negotiating deal...
💌  Deal agreed. Running job...
∙∙∙ Job submitted. Negotiating deal...Deal QmQZgZDo6YJFrBZNUfQhdK2zYdFB47KJgcPDnEBqH3poNg is matched
∙●∙ Job submitted. Negotiating deal...2024/03/08 07:21:53 Deal QmQZgZDo6YJFrBZNUfQhdK2zYdFB47KJgcPDnEBqH3poNg is agreed
🤝  Job submitted. Negotiating deal...
💌  Deal agreed. Running job...
🤔  Results submitted. Awaiting verification...
✅  Results accepted. Downloading result...

🍂 CoopHive job completed in 58.219397459s, try 👇
    open /Users/hiro/.coophive/sepolia/jc/downloaded-files/QmSirREjZ9icDawWiR8EMQn8UaVMFySrocgfdGstCrjiGR
    cat /Users/hiro/.coophive/sepolia/jc/downloaded-files/QmSirREjZ9icDawWiR8EMQn8UaVMFySrocgfdGstCrjiGR/stdout
    cat /Users/hiro/.coophive/sepolia/jc/downloaded-files/QmSirREjZ9icDawWiR8EMQn8UaVMFySrocgfdGstCrjiGR/stderr
    https://ipfs.io/ipfs/QmRnfzKiR2hmm5UeSx2Cn2CRFEB1gTDBXavf4oiJC1Aeoo
</pre>


<pre>
cat /Users/hiro/.coophive/sepolia/jc/downloaded-files/QmSirREjZ9icDawWiR8EMQn8UaVMFySrocgfdGstCrjiGR/stdout 
 ___________________________
< Hiro is building CoopHive >
 ---------------------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
</pre>

</details>

<details> 
<summary>Cowsay on Halcyon</summary>

https://github.com/CoopHive/hive/assets/24226219/4edb669c-7eb7-4dee-9adf-1651ddb73b3b

</details>

### SDXL v0.9.0

Stable diffusion:

```
hive run sdxl:v0.3.0 -i Prompt="hiro saves the hive"
```


https://github.com/CoopHive/hive/assets/24226219/b0a00864-5afb-49ee-9bb6-a44a8936194e


<details>

<summary> Preview Run </summary>
<pre>
hive run sdxl:v0.3.0 -i Prompt="hiro saves the hive"

  ___  __    __  ____  _  _  __  _  _  ____
 / __)/  \  /  \(  _ \/ )( \(  )/ )( \(  __)
( (__(  O )(  O )) __/) __ ( )( \ \/ / ) _)
 \___)\__/  \__/(__)  \_)(_/(__) \__/ (____) 0.17.1

  Decentralized Compute Network  https://coophive.network


∙●∙ CoopHive submitting job2024-03-08T07:22:41+05:30 INF ../../Volumes/N/CoopHive/hive/pkg/module/utils.go:148 > updating cached git repo=/Users/hiro/.coophive/sepolia/jc/repos/CoopHive/coophive-module-sdxl
🌟  CoopHive submitting job
∙∙∙ Job submitted. Negotiating deal...Deal QmbBhg8NqYyqYp6QnUhd1yZgYhRoKYt1bZLxKtUjkLNEbM is matched
●∙∙ Job submitted. Negotiating deal...2024/03/08 07:22:43 Deal QmbBhg8NqYyqYp6QnUhd1yZgYhRoKYt1bZLxKtUjkLNEbM is agreed
🤝  Job submitted. Negotiating deal...
💌  Deal agreed. Running job...
🤔  Results submitted. Awaiting verification...
✅  Results accepted. Downloading result...
🤔  Results submitted. Awaiting verification...
✅  Results accepted. Downloading result...

🍂 CoopHive job completed in 2m43.038262916s, try 👇
    open /Users/hiro/.coophive/sepolia/jc/downloaded-files/QmbBhg8NqYyqYp6QnUhd1yZgYhRoKYt1bZLxKtUjkLNEbM
    cat /Users/hiro/.coophive/sepolia/jc/downloaded-files/QmbBhg8NqYyqYp6QnUhd1yZgYhRoKYt1bZLxKtUjkLNEbM/stdout
    cat /Users/hiro/.coophive/sepolia/jc/downloaded-files/QmbBhg8NqYyqYp6QnUhd1yZgYhRoKYt1bZLxKtUjkLNEbM/stderr
    https://ipfs.io/ipfs/Qmf8jBCQHz6io6VK8GnQZVzk7hUHG6NWXjqQs4AeEervEv

</pre>

<pre>
cat /Users/hiro/.coophive/sepolia/jc/downloaded-files/QmbBhg8NqYyqYp6QnUhd1yZgYhRoKYt1bZLxKtUjkLNEbM/stdout
Random seed set as 12
Got 1 images
</pre>

![SDXL:v0.3.0](https://github.com/CoopHive/hive/assets/24226219/92c3c52f-9b3a-49c2-8aa1-7d28d87d9709)

> Didn't like the image? Try a different seed
>> hive run sdxl:v0.3.0 -i Prompt="hiro saves the hive" -i Seed=16

</details>

[//]: # (## Run a node, earn HIVE)

[//]: # ()

[//]: # ()

[//]: # (```)

[//]: # (hive rp)

[//]: # (```)

[//]: # ()

[//]: # (Deploy seamlessly on linux by utilizing [these systemd configuration files]&#40;https://github.com/CoopHive/hive/tree/main/ops&#41;.)

## Available modules

Check the github releases page for each module or just use the git hash as the tag.

- [sdxl](https://github.com/CoopHive/coophive-module-sdxl)
- [stable-diffusion](https://github.com/CoopHive/coophive-module-stable-diffusion)
- [duckdb](https://github.com/CoopHive/coophive-module-duckdb)
- [fastchat](https://github.com/CoopHive/coophive-module-fastchat)
- [lora-inference](https://github.com/CoopHive/coophive-module-lora-inference)
- [lora-training](https://github.com/CoopHive/coophive-module-lora-training)
- [filecoin-data-prep](https://github.com/CoopHive/coophive-module-filecoin-data-prep)
- [wasm](https://github.com/CoopHive/coophive-module-wasm)
- [cowsay](https://github.com/CoopHive/coophive-module-cowsay)

## Write a module

A module is just a git repo, and module versions are just git tags.

In your repo, create a file called `module.coophive`. For an example, see [cowsay](https://github.com/CoopHive/coophive-module-cowsay).

This is a json template with Go text/template style `{{.Message}}` sections which will be replaced by CoopHive with json
encoded inputs to modules. You can also do fancy things with go templates like setting defaults, see cowsay for example.
While developing a module, you can use the git hash to test it.

Pass inputs as:

```
hive run github.com/username/repo:tag -i InputVar=value
```

Inputs are a map of strings to strings.

### Writing Advanced Modules

1. `subt`:
   The `subt` function allows for substitutions in your template.

This function is a workaround for the lack of direct substitution support in the module. It implements
the [printf](https://pkg.go.dev/text/template#Template.Funcs) function under the hood, which allows you to format
strings with placeholders.

<details>
  <summary> 
    Usage   
  </summary>
    The `subt` function can be used in the same way as the `printf` function in Go. You pass in a format string, followed by values that correspond to the placeholders in the format string.
    ```
    const templateText = `
    {{ subt "Hello %s" .name }}
    `
    ```
</details>
[Sample Module for Ref](https://github.com/CoopHive/coophive-module-marker/blob/243e86c/module.coophive#L17)
[Prototype Your Module over here](https://go.dev/play/p/oBgc2Cetug3)

2. `subst`: same as `subt`. Checkout https://github.com/CoopHive/coophive-module-sdxl/blob/v0.3.0/module.coophive#L24

3. `or`: takes 2 positional arguments : inputA, inputB. It returns the inputA if its a non-empty string else inputB. It
   can be used to ensure default vars without nested conditional
   logics: https://github.com/CoopHive/coophive-module-marker/blob/243e86c/module.coophive#L47. You can prototype your
   template strings using this scratch file: https://go.dev/play/p/_FvTYbtKim1

4. `get`: alias of `or`

> If you are golang dev, you can better insights on implementation of the template functions by checking
> out: <a href=./pkg/module/templateFuncs_test.go> Template Func Test</a>

[CoopHive]: https://coophive.network

[Aurora RPC]: http://aurora.co-ophive.network:8545

[Aurora Faucet]: http://faucet.co-ophive.network:8080

[Halcyon RPC]: http://halcyon.co-ophive.network:8545

[Halcyon Faucet]: http://halcyon-faucet.co-ophive.network:8085

