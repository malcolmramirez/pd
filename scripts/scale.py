import argparse
import pprint

cents_per_octave = 1200
reference_freq = 440 # A4
base_octave = 4

def freq(cents: int):
    return reference_freq * (2 ** (cents / cents_per_octave))

def edo(divisions: int, octaves: tuple):
    assert divisions > 0

    start_octave, end_octave = octaves
    cents_per_note = cents_per_octave // divisions

    octaves_to_cents = [
        (octave, (octave - base_octave) * cents_per_octave)
        for octave in range(start_octave, end_octave + 1)
    ]

    return {
        octave: [freq(octave_cents + div) for div in range(0, cents_per_octave, cents_per_note)]
        for octave, octave_cents in octaves_to_cents
    }


def main():
    parser = argparse.ArgumentParser(
        prog='Note frequency generator',
        description='Note frequency generator')
    parser.add_argument('-d', '--edo', type=int)
    parser.add_argument('-s', '--start-octave', type=int)
    parser.add_argument('-e', '--end-octave', type=int)

    args = parser.parse_args()

    octave_range = (args.start_octave, args.end_octave)
    table = edo(args.edo, octave_range)
    pprint.pp(table)

if __name__ == "__main__":
    main()
